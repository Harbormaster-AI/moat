
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { System_Service } from '../../../services/System_.service';
import { System_ } from '../../../models/System_';

@Component({
    selector: 'app-index-system_',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexSystem_Component implements OnInit {

    system_s: System_[] = [];

    constructor(
        private router: Router,
        private service: System_Service
) {}

    ngOnInit(): void {
        this.getSystem_s();
}

    getSystem_s(): void {
        this.service.getSystem_s().subscribe((res) => {
        this.system_s = res;
    });
}

    deleteSystem_(id: any): void {
        this.service.deleteSystem_(id)
            .subscribe(() => {
                this.getSystem_s();
            });
    }
}