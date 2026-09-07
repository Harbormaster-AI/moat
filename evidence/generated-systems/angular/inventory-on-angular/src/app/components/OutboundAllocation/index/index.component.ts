
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { OutboundAllocationService } from '../../../services/OutboundAllocation.service';
import { OutboundAllocation } from '../../../models/OutboundAllocation';

@Component({
    selector: 'app-index-outboundAllocation',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexOutboundAllocationComponent implements OnInit {

    outboundAllocations: OutboundAllocation[] = [];

    constructor(
        private router: Router,
        private service: OutboundAllocationService
) {}

    ngOnInit(): void {
        this.getOutboundAllocations();
}

    getOutboundAllocations(): void {
        this.service.getOutboundAllocations().subscribe((res) => {
        this.outboundAllocations = res;
    });
}

    deleteOutboundAllocation(id: any): void {
        this.service.deleteOutboundAllocation(id)
            .subscribe(() => {
                this.getOutboundAllocations();
            });
    }
}