
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LandingGearService } from '../../../services/LandingGear.service';
import { LandingGear } from '../../../models/LandingGear';

@Component({
    selector: 'app-index-landingGear',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexLandingGearComponent implements OnInit {

    landingGears: LandingGear[] = [];

    constructor(
        private router: Router,
        private service: LandingGearService
) {}

    ngOnInit(): void {
        this.getLandingGears();
}

    getLandingGears(): void {
        this.service.getLandingGears().subscribe((res) => {
        this.landingGears = res;
    });
}

    deleteLandingGear(id: any): void {
        this.service.deleteLandingGear(id)
            .subscribe(() => {
                this.getLandingGears();
            });
    }
}