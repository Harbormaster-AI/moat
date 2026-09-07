
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Model_Service } from '../../../services/Model_.service';
import { Model_ } from '../../../models/Model_';

@Component({
    selector: 'app-index-model_',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexModel_Component implements OnInit {

    model_s: Model_[] = [];

    constructor(
        private router: Router,
        private service: Model_Service
) {}

    ngOnInit(): void {
        this.getModel_s();
}

    getModel_s(): void {
        this.service.getModel_s().subscribe((res) => {
        this.model_s = res;
    });
}

    deleteModel_(id: any): void {
        this.service.deleteModel_(id)
            .subscribe(() => {
                this.getModel_s();
            });
    }
}