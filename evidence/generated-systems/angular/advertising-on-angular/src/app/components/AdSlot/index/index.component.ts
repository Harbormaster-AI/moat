
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AdSlotService } from '../../../services/AdSlot.service';
import { AdSlot } from '../../../models/AdSlot';

@Component({
    selector: 'app-index-adSlot',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAdSlotComponent implements OnInit {

    adSlots: AdSlot[] = [];

    constructor(
        private router: Router,
        private service: AdSlotService
) {}

    ngOnInit(): void {
        this.getAdSlots();
}

    getAdSlots(): void {
        this.service.getAdSlots().subscribe((res) => {
        this.adSlots = res;
    });
}

    deleteAdSlot(id: any): void {
        this.service.deleteAdSlot(id)
            .subscribe(() => {
                this.getAdSlots();
            });
    }
}