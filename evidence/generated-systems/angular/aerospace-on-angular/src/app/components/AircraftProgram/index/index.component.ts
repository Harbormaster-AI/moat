
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AircraftProgramService } from '../../../services/AircraftProgram.service';
import { AircraftProgram } from '../../../models/AircraftProgram';

@Component({
    selector: 'app-index-aircraftProgram',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAircraftProgramComponent implements OnInit {

    aircraftPrograms: AircraftProgram[] = [];

    constructor(
        private router: Router,
        private service: AircraftProgramService
) {}

    ngOnInit(): void {
        this.getAircraftPrograms();
}

    getAircraftPrograms(): void {
        this.service.getAircraftPrograms().subscribe((res) => {
        this.aircraftPrograms = res;
    });
}

    deleteAircraftProgram(id: any): void {
        this.service.deleteAircraftProgram(id)
            .subscribe(() => {
                this.getAircraftPrograms();
            });
    }
}