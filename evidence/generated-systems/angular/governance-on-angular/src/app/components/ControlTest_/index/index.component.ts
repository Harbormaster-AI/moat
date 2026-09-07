
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ControlTest_Service } from '../../../services/ControlTest_.service';
import { ControlTest_ } from '../../../models/ControlTest_';

@Component({
    selector: 'app-index-controlTest_',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexControlTest_Component implements OnInit {

    controlTest_s: ControlTest_[] = [];

    constructor(
        private router: Router,
        private service: ControlTest_Service
) {}

    ngOnInit(): void {
        this.getControlTest_s();
}

    getControlTest_s(): void {
        this.service.getControlTest_s().subscribe((res) => {
        this.controlTest_s = res;
    });
}

    deleteControlTest_(id: any): void {
        this.service.deleteControlTest_(id)
            .subscribe(() => {
                this.getControlTest_s();
            });
    }
}