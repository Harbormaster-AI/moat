
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AircraftFamilyService } from '../../../services/AircraftFamily.service';
import { AircraftFamily } from '../../../models/AircraftFamily';

@Component({
    selector: 'app-index-aircraftFamily',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAircraftFamilyComponent implements OnInit {

    aircraftFamilys: AircraftFamily[] = [];

    constructor(
        private router: Router,
        private service: AircraftFamilyService
) {}

    ngOnInit(): void {
        this.getAircraftFamilys();
}

    getAircraftFamilys(): void {
        this.service.getAircraftFamilys().subscribe((res) => {
        this.aircraftFamilys = res;
    });
}

    deleteAircraftFamily(id: any): void {
        this.service.deleteAircraftFamily(id)
            .subscribe(() => {
                this.getAircraftFamilys();
            });
    }
}