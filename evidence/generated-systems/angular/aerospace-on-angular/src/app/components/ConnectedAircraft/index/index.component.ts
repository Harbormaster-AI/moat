
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ConnectedAircraftService } from '../../../services/ConnectedAircraft.service';
import { ConnectedAircraft } from '../../../models/ConnectedAircraft';

@Component({
    selector: 'app-index-connectedAircraft',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexConnectedAircraftComponent implements OnInit {

    connectedAircrafts: ConnectedAircraft[] = [];

    constructor(
        private router: Router,
        private service: ConnectedAircraftService
) {}

    ngOnInit(): void {
        this.getConnectedAircrafts();
}

    getConnectedAircrafts(): void {
        this.service.getConnectedAircrafts().subscribe((res) => {
        this.connectedAircrafts = res;
    });
}

    deleteConnectedAircraft(id: any): void {
        this.service.deleteConnectedAircraft(id)
            .subscribe(() => {
                this.getConnectedAircrafts();
            });
    }
}