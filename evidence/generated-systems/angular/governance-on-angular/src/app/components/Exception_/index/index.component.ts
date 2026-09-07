
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Exception_Service } from '../../../services/Exception_.service';
import { Exception_ } from '../../../models/Exception_';

@Component({
    selector: 'app-index-exception_',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexException_Component implements OnInit {

    exception_s: Exception_[] = [];

    constructor(
        private router: Router,
        private service: Exception_Service
) {}

    ngOnInit(): void {
        this.getException_s();
}

    getException_s(): void {
        this.service.getException_s().subscribe((res) => {
        this.exception_s = res;
    });
}

    deleteException_(id: any): void {
        this.service.deleteException_(id)
            .subscribe(() => {
                this.getException_s();
            });
    }
}