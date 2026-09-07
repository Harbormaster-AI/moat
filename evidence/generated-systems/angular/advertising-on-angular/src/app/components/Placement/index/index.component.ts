
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PlacementService } from '../../../services/Placement.service';
import { Placement } from '../../../models/Placement';

@Component({
    selector: 'app-index-placement',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPlacementComponent implements OnInit {

    placements: Placement[] = [];

    constructor(
        private router: Router,
        private service: PlacementService
) {}

    ngOnInit(): void {
        this.getPlacements();
}

    getPlacements(): void {
        this.service.getPlacements().subscribe((res) => {
        this.placements = res;
    });
}

    deletePlacement(id: any): void {
        this.service.deletePlacement(id)
            .subscribe(() => {
                this.getPlacements();
            });
    }
}