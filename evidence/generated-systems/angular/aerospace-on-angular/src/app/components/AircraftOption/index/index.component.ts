
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AircraftOptionService } from '../../../services/AircraftOption.service';
import { AircraftOption } from '../../../models/AircraftOption';

@Component({
    selector: 'app-index-aircraftOption',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAircraftOptionComponent implements OnInit {

    aircraftOptions: AircraftOption[] = [];

    constructor(
        private router: Router,
        private service: AircraftOptionService
) {}

    ngOnInit(): void {
        this.getAircraftOptions();
}

    getAircraftOptions(): void {
        this.service.getAircraftOptions().subscribe((res) => {
        this.aircraftOptions = res;
    });
}

    deleteAircraftOption(id: any): void {
        this.service.deleteAircraftOption(id)
            .subscribe(() => {
                this.getAircraftOptions();
            });
    }
}