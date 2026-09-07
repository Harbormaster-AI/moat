
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LocationService } from '../../../services/Location.service';
import { Location } from '../../../models/Location';

@Component({
    selector: 'app-index-location',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexLocationComponent implements OnInit {

    locations: Location[] = [];

    constructor(
        private router: Router,
        private service: LocationService
) {}

    ngOnInit(): void {
        this.getLocations();
}

    getLocations(): void {
        this.service.getLocations().subscribe((res) => {
        this.locations = res;
    });
}

    deleteLocation(id: any): void {
        this.service.deleteLocation(id)
            .subscribe(() => {
                this.getLocations();
            });
    }
}