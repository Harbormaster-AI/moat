
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InsuredObjectService } from '../../../services/InsuredObject.service';
import { InsuredObject } from '../../../models/InsuredObject';

@Component({
    selector: 'app-index-insuredObject',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInsuredObjectComponent implements OnInit {

    insuredObjects: InsuredObject[] = [];

    constructor(
        private router: Router,
        private service: InsuredObjectService
) {}

    ngOnInit(): void {
        this.getInsuredObjects();
}

    getInsuredObjects(): void {
        this.service.getInsuredObjects().subscribe((res) => {
        this.insuredObjects = res;
    });
}

    deleteInsuredObject(id: any): void {
        this.service.deleteInsuredObject(id)
            .subscribe(() => {
                this.getInsuredObjects();
            });
    }
}