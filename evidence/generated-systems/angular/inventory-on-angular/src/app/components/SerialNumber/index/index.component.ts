
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SerialNumberService } from '../../../services/SerialNumber.service';
import { SerialNumber } from '../../../models/SerialNumber';

@Component({
    selector: 'app-index-serialNumber',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexSerialNumberComponent implements OnInit {

    serialNumbers: SerialNumber[] = [];

    constructor(
        private router: Router,
        private service: SerialNumberService
) {}

    ngOnInit(): void {
        this.getSerialNumbers();
}

    getSerialNumbers(): void {
        this.service.getSerialNumbers().subscribe((res) => {
        this.serialNumbers = res;
    });
}

    deleteSerialNumber(id: any): void {
        this.service.deleteSerialNumber(id)
            .subscribe(() => {
                this.getSerialNumbers();
            });
    }
}