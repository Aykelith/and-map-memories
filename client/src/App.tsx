import { Component, createResource } from 'solid-js';

import { createEffect, onMount, createSignal } from 'solid-js';

import 'leaflet/dist/leaflet.css';
import * as L from 'leaflet';

import 'leaflet.markercluster';

import classNames from "classnames";

import markerIcon from "../node_modules/leaflet/dist/images/marker-icon.png";
import { getStoreItem, setStoreItem } from './utils/store';
L.Marker.prototype.setIcon(L.icon({
    iconUrl:markerIcon,
    iconSize: [24,36],
    iconAnchor: [12,36]
}));

const Dialog: Component<any> = (props) => {
    return <div class={classNames("absolute top-0 left-0 w-full h-full flex justify-center items-center pointer-events-none z-[1000] bg-neutral-950/50 lg:bg-inherit", props.class)}>
        {props.children}
    </div>
};

const PinOverflow: Component<any> = (props) => {
    const [isLoading, setLoading] = createSignal<boolean>(true);

    onMount(() => {

    });

    return (
        <div class="relative w-11/12 lg:w-4/12 h-5/6 lg:h-full bg-stone-50 shadow rounded lg:rounded-none pointer-events-auto pt-4 pb-2 px-2">
            <button class="absolute top-0 right-0 px-1.5 pb-0.5 font-bold" onClick={props.onClose}>✕</button>
            <div class="font-bold">Grădina Botanică „Vasile Fati” din Jibou</div>
        </div>
    );
};

async function fetchPins() {
    const response = await fetch("/api/pins/get/all");
    const jsonData = await response.json();
    return jsonData;
}

async function loadPins(onNewPin) {
    const pins = await fetchPins();
    const actualExistingPins = getStoreItem("pins") || [];

    const newPins = [];
    for (const pin of pins) {
        let found = false;
        for (const existingPin of actualExistingPins) {
            if (pin.ID === existingPin.ID) {
                found = true;
                break;
            }
        }

        if (!found) newPins.push(pin);
    }

    const map = getStoreItem("map");

    let pointClusters = L.markerClusterGroup();
    newPins.map(pinData => {
        const marker = L.marker([pinData.Lat, pinData.Lng]);
        onNewPin(pinData, marker);
        pointClusters.addLayer(marker); 
    });
    map.addLayer(pointClusters);
}

const App: Component = () => {
    const [currentDisplayedPinID, setCurrentDisplayedPinID] = createSignal<string | null>(null);

    const [isLoading, setLoading] = createSignal<boolean>(true);

    onMount(() => {
        const map = L.map('map').setView([45.714, 24.873], 7);
        
        L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
            attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        }).addTo(map);

        setStoreItem("map", map);

        loadPins((pinData, marker) => {
            marker.on('click', () => setCurrentDisplayedPinID(pinData.ID));
        });
    });
    

    return (
        <div class="w-full h-full">
        <div id="map"></div>
            { currentDisplayedPinID() !== null && <Dialog class="lg:justify-end"><PinOverflow pinID={currentDisplayedPinID()} onClose={() => setCurrentDisplayedPinID(null)}/></Dialog> }
        </div>
    );
};

export default App;
