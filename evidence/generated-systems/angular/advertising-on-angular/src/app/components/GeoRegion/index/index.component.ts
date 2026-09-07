
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { GeoRegionService } from '../../../services/GeoRegion.service';
import { GeoRegion } from '../../../models/GeoRegion';

@Component({
    selector: 'app-index-geoRegion',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexGeoRegionComponent implements OnInit {

    geoRegions: GeoRegion[] = [];

    constructor(
        private router: Router,
        private service: GeoRegionService
) {}

    ngOnInit(): void {
        this.getGeoRegions();
}

    getGeoRegions(): void {
        this.service.getGeoRegions().subscribe((res) => {
        this.geoRegions = res;
    });
}

    deleteGeoRegion(id: any): void {
        this.service.deleteGeoRegion(id)
            .subscribe(() => {
                this.getGeoRegions();
            });
    }
}