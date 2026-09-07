
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ImagingCenterService } from '../../../services/ImagingCenter.service';
import { ImagingCenter } from '../../../models/ImagingCenter';

@Component({
    selector: 'app-index-imagingCenter',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexImagingCenterComponent implements OnInit {

    imagingCenters: ImagingCenter[] = [];

    constructor(
        private router: Router,
        private service: ImagingCenterService
) {}

    ngOnInit(): void {
        this.getImagingCenters();
}

    getImagingCenters(): void {
        this.service.getImagingCenters().subscribe((res) => {
        this.imagingCenters = res;
    });
}

    deleteImagingCenter(id: any): void {
        this.service.deleteImagingCenter(id)
            .subscribe(() => {
                this.getImagingCenters();
            });
    }
}