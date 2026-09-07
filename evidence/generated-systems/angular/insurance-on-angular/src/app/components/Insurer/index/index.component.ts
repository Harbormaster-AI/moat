
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InsurerService } from '../../../services/Insurer.service';
import { Insurer } from '../../../models/Insurer';

@Component({
    selector: 'app-index-insurer',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInsurerComponent implements OnInit {

    insurers: Insurer[] = [];

    constructor(
        private router: Router,
        private service: InsurerService
) {}

    ngOnInit(): void {
        this.getInsurers();
}

    getInsurers(): void {
        this.service.getInsurers().subscribe((res) => {
        this.insurers = res;
    });
}

    deleteInsurer(id: any): void {
        this.service.deleteInsurer(id)
            .subscribe(() => {
                this.getInsurers();
            });
    }
}