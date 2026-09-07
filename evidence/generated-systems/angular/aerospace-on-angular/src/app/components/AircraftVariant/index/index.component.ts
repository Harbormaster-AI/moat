
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AircraftVariantService } from '../../../services/AircraftVariant.service';
import { AircraftVariant } from '../../../models/AircraftVariant';

@Component({
    selector: 'app-index-aircraftVariant',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAircraftVariantComponent implements OnInit {

    aircraftVariants: AircraftVariant[] = [];

    constructor(
        private router: Router,
        private service: AircraftVariantService
) {}

    ngOnInit(): void {
        this.getAircraftVariants();
}

    getAircraftVariants(): void {
        this.service.getAircraftVariants().subscribe((res) => {
        this.aircraftVariants = res;
    });
}

    deleteAircraftVariant(id: any): void {
        this.service.deleteAircraftVariant(id)
            .subscribe(() => {
                this.getAircraftVariants();
            });
    }
}