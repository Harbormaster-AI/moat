
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SubrogationRecoveryService } from '../../../services/SubrogationRecovery.service';
import { SubrogationRecovery } from '../../../models/SubrogationRecovery';

@Component({
    selector: 'app-index-subrogationRecovery',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexSubrogationRecoveryComponent implements OnInit {

    subrogationRecoverys: SubrogationRecovery[] = [];

    constructor(
        private router: Router,
        private service: SubrogationRecoveryService
) {}

    ngOnInit(): void {
        this.getSubrogationRecoverys();
}

    getSubrogationRecoverys(): void {
        this.service.getSubrogationRecoverys().subscribe((res) => {
        this.subrogationRecoverys = res;
    });
}

    deleteSubrogationRecovery(id: any): void {
        this.service.deleteSubrogationRecovery(id)
            .subscribe(() => {
                this.getSubrogationRecoverys();
            });
    }
}