
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BusinessUnitService } from '../../../services/BusinessUnit.service';
import { BusinessUnit } from '../../../models/BusinessUnit';

@Component({
    selector: 'app-index-businessUnit',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexBusinessUnitComponent implements OnInit {

    businessUnits: BusinessUnit[] = [];

    constructor(
        private router: Router,
        private service: BusinessUnitService
) {}

    ngOnInit(): void {
        this.getBusinessUnits();
}

    getBusinessUnits(): void {
        this.service.getBusinessUnits().subscribe((res) => {
        this.businessUnits = res;
    });
}

    deleteBusinessUnit(id: any): void {
        this.service.deleteBusinessUnit(id)
            .subscribe(() => {
                this.getBusinessUnits();
            });
    }
}