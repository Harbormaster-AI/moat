
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DSPService } from '../../../services/DSP.service';
import { DSP } from '../../../models/DSP';

@Component({
    selector: 'app-index-dSP',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDSPComponent implements OnInit {

    dSPs: DSP[] = [];

    constructor(
        private router: Router,
        private service: DSPService
) {}

    ngOnInit(): void {
        this.getDSPs();
}

    getDSPs(): void {
        this.service.getDSPs().subscribe((res) => {
        this.dSPs = res;
    });
}

    deleteDSP(id: any): void {
        this.service.deleteDSP(id)
            .subscribe(() => {
                this.getDSPs();
            });
    }
}