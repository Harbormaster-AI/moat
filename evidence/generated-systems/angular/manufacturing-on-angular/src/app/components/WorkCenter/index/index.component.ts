
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { WorkCenterService } from '../../../services/WorkCenter.service';
import { WorkCenter } from '../../../models/WorkCenter';

@Component({
    selector: 'app-index-workCenter',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexWorkCenterComponent implements OnInit {

    workCenters: WorkCenter[] = [];

    constructor(
        private router: Router,
        private service: WorkCenterService
) {}

    ngOnInit(): void {
        this.getWorkCenters();
}

    getWorkCenters(): void {
        this.service.getWorkCenters().subscribe((res) => {
        this.workCenters = res;
    });
}

    deleteWorkCenter(id: any): void {
        this.service.deleteWorkCenter(id)
            .subscribe(() => {
                this.getWorkCenters();
            });
    }
}