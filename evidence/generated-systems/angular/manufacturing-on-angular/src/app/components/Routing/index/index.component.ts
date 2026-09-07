
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RoutingService } from '../../../services/Routing.service';
import { Routing } from '../../../models/Routing';

@Component({
    selector: 'app-index-routing',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexRoutingComponent implements OnInit {

    routings: Routing[] = [];

    constructor(
        private router: Router,
        private service: RoutingService
) {}

    ngOnInit(): void {
        this.getRoutings();
}

    getRoutings(): void {
        this.service.getRoutings().subscribe((res) => {
        this.routings = res;
    });
}

    deleteRouting(id: any): void {
        this.service.deleteRouting(id)
            .subscribe(() => {
                this.getRoutings();
            });
    }
}