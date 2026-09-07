
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AircraftModelService } from '../../../services/AircraftModel.service';
import { AircraftModel } from '../../../models/AircraftModel';

@Component({
    selector: 'app-index-aircraftModel',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAircraftModelComponent implements OnInit {

    aircraftModels: AircraftModel[] = [];

    constructor(
        private router: Router,
        private service: AircraftModelService
) {}

    ngOnInit(): void {
        this.getAircraftModels();
}

    getAircraftModels(): void {
        this.service.getAircraftModels().subscribe((res) => {
        this.aircraftModels = res;
    });
}

    deleteAircraftModel(id: any): void {
        this.service.deleteAircraftModel(id)
            .subscribe(() => {
                this.getAircraftModels();
            });
    }
}