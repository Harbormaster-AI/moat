
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Case_Service } from '../../../services/Case_.service';
import { Case_ } from '../../../models/Case_';

@Component({
    selector: 'app-index-case_',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCase_Component implements OnInit {

    case_s: Case_[] = [];

    constructor(
        private router: Router,
        private service: Case_Service
) {}

    ngOnInit(): void {
        this.getCase_s();
}

    getCase_s(): void {
        this.service.getCase_s().subscribe((res) => {
        this.case_s = res;
    });
}

    deleteCase_(id: any): void {
        this.service.deleteCase_(id)
            .subscribe(() => {
                this.getCase_s();
            });
    }
}