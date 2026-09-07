
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { QuarantineService } from '../../../services/Quarantine.service';
import { Quarantine } from '../../../models/Quarantine';

@Component({
    selector: 'app-index-quarantine',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexQuarantineComponent implements OnInit {

    quarantines: Quarantine[] = [];

    constructor(
        private router: Router,
        private service: QuarantineService
) {}

    ngOnInit(): void {
        this.getQuarantines();
}

    getQuarantines(): void {
        this.service.getQuarantines().subscribe((res) => {
        this.quarantines = res;
    });
}

    deleteQuarantine(id: any): void {
        this.service.deleteQuarantine(id)
            .subscribe(() => {
                this.getQuarantines();
            });
    }
}