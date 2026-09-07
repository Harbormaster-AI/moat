
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AircraftPackageService } from '../../../services/AircraftPackage.service';
import { AircraftPackage } from '../../../models/AircraftPackage';

@Component({
    selector: 'app-index-aircraftPackage',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAircraftPackageComponent implements OnInit {

    aircraftPackages: AircraftPackage[] = [];

    constructor(
        private router: Router,
        private service: AircraftPackageService
) {}

    ngOnInit(): void {
        this.getAircraftPackages();
}

    getAircraftPackages(): void {
        this.service.getAircraftPackages().subscribe((res) => {
        this.aircraftPackages = res;
    });
}

    deleteAircraftPackage(id: any): void {
        this.service.deleteAircraftPackage(id)
            .subscribe(() => {
                this.getAircraftPackages();
            });
    }
}