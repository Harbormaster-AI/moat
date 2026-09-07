
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ExposureService } from '../../../services/Exposure.service';
import { Exposure } from '../../../models/Exposure';

@Component({
    selector: 'app-index-exposure',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexExposureComponent implements OnInit {

    exposures: Exposure[] = [];

    constructor(
        private router: Router,
        private service: ExposureService
) {}

    ngOnInit(): void {
        this.getExposures();
}

    getExposures(): void {
        this.service.getExposures().subscribe((res) => {
        this.exposures = res;
    });
}

    deleteExposure(id: any): void {
        this.service.deleteExposure(id)
            .subscribe(() => {
                this.getExposures();
            });
    }
}