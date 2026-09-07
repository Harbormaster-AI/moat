
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Record_Service } from '../../../services/Record_.service';
import { Record_ } from '../../../models/Record_';

@Component({
    selector: 'app-index-record_',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexRecord_Component implements OnInit {

    record_s: Record_[] = [];

    constructor(
        private router: Router,
        private service: Record_Service
) {}

    ngOnInit(): void {
        this.getRecord_s();
}

    getRecord_s(): void {
        this.service.getRecord_s().subscribe((res) => {
        this.record_s = res;
    });
}

    deleteRecord_(id: any): void {
        this.service.deleteRecord_(id)
            .subscribe(() => {
                this.getRecord_s();
            });
    }
}