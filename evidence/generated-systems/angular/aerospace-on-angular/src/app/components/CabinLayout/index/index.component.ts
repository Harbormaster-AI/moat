
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CabinLayoutService } from '../../../services/CabinLayout.service';
import { CabinLayout } from '../../../models/CabinLayout';

@Component({
    selector: 'app-index-cabinLayout',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCabinLayoutComponent implements OnInit {

    cabinLayouts: CabinLayout[] = [];

    constructor(
        private router: Router,
        private service: CabinLayoutService
) {}

    ngOnInit(): void {
        this.getCabinLayouts();
}

    getCabinLayouts(): void {
        this.service.getCabinLayouts().subscribe((res) => {
        this.cabinLayouts = res;
    });
}

    deleteCabinLayout(id: any): void {
        this.service.deleteCabinLayout(id)
            .subscribe(() => {
                this.getCabinLayouts();
            });
    }
}