
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UoMConversionService } from '../../../services/UoMConversion.service';
import { UoMConversion } from '../../../models/UoMConversion';

@Component({
    selector: 'app-index-uoMConversion',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexUoMConversionComponent implements OnInit {

    uoMConversions: UoMConversion[] = [];

    constructor(
        private router: Router,
        private service: UoMConversionService
) {}

    ngOnInit(): void {
        this.getUoMConversions();
}

    getUoMConversions(): void {
        this.service.getUoMConversions().subscribe((res) => {
        this.uoMConversions = res;
    });
}

    deleteUoMConversion(id: any): void {
        this.service.deleteUoMConversion(id)
            .subscribe(() => {
                this.getUoMConversions();
            });
    }
}