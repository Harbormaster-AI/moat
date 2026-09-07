
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FacilityService } from '../../../services/Facility.service';
import { Facility } from '../../../models/Facility';

@Component({
    selector: 'app-index-facility',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexFacilityComponent implements OnInit {

    facilitys: Facility[] = [];

    constructor(
        private router: Router,
        private service: FacilityService
) {}

    ngOnInit(): void {
        this.getFacilitys();
}

    getFacilitys(): void {
        this.service.getFacilitys().subscribe((res) => {
        this.facilitys = res;
    });
}

    deleteFacility(id: any): void {
        this.service.deleteFacility(id)
            .subscribe(() => {
                this.getFacilitys();
            });
    }
}