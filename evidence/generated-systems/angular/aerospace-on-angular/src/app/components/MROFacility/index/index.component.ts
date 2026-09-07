
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MROFacilityService } from '../../../services/MROFacility.service';
import { MROFacility } from '../../../models/MROFacility';

@Component({
    selector: 'app-index-mROFacility',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexMROFacilityComponent implements OnInit {

    mROFacilitys: MROFacility[] = [];

    constructor(
        private router: Router,
        private service: MROFacilityService
) {}

    ngOnInit(): void {
        this.getMROFacilitys();
}

    getMROFacilitys(): void {
        this.service.getMROFacilitys().subscribe((res) => {
        this.mROFacilitys = res;
    });
}

    deleteMROFacility(id: any): void {
        this.service.deleteMROFacility(id)
            .subscribe(() => {
                this.getMROFacilitys();
            });
    }
}