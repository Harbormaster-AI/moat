
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AircraftService } from '../../../services/Aircraft.service';
import { Aircraft } from '../../../models/Aircraft';

@Component({
    selector: 'app-index-aircraft',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAircraftComponent implements OnInit {

    aircrafts: Aircraft[] = [];

    constructor(
        private router: Router,
        private service: AircraftService
) {}

    ngOnInit(): void {
        this.getAircrafts();
}

    getAircrafts(): void {
        this.service.getAircrafts().subscribe((res) => {
        this.aircrafts = res;
    });
}

    deleteAircraft(id: any): void {
        this.service.deleteAircraft(id)
            .subscribe(() => {
                this.getAircrafts();
            });
    }
}