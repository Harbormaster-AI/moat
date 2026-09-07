
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DeviceCriterionService } from '../../../services/DeviceCriterion.service';
import { DeviceCriterion } from '../../../models/DeviceCriterion';

@Component({
    selector: 'app-index-deviceCriterion',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDeviceCriterionComponent implements OnInit {

    deviceCriterions: DeviceCriterion[] = [];

    constructor(
        private router: Router,
        private service: DeviceCriterionService
) {}

    ngOnInit(): void {
        this.getDeviceCriterions();
}

    getDeviceCriterions(): void {
        this.service.getDeviceCriterions().subscribe((res) => {
        this.deviceCriterions = res;
    });
}

    deleteDeviceCriterion(id: any): void {
        this.service.deleteDeviceCriterion(id)
            .subscribe(() => {
                this.getDeviceCriterions();
            });
    }
}