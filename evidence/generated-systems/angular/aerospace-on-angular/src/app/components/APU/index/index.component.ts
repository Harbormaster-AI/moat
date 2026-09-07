
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { APUService } from '../../../services/APU.service';
import { APU } from '../../../models/APU';

@Component({
    selector: 'app-index-aPU',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAPUComponent implements OnInit {

    aPUs: APU[] = [];

    constructor(
        private router: Router,
        private service: APUService
) {}

    ngOnInit(): void {
        this.getAPUs();
}

    getAPUs(): void {
        this.service.getAPUs().subscribe((res) => {
        this.aPUs = res;
    });
}

    deleteAPU(id: any): void {
        this.service.deleteAPU(id)
            .subscribe(() => {
                this.getAPUs();
            });
    }
}