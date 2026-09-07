
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DimensionService } from '../../../services/Dimension.service';
import { Dimension } from '../../../models/Dimension';

@Component({
    selector: 'app-index-dimension',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDimensionComponent implements OnInit {

    dimensions: Dimension[] = [];

    constructor(
        private router: Router,
        private service: DimensionService
) {}

    ngOnInit(): void {
        this.getDimensions();
}

    getDimensions(): void {
        this.service.getDimensions().subscribe((res) => {
        this.dimensions = res;
    });
}

    deleteDimension(id: any): void {
        this.service.deleteDimension(id)
            .subscribe(() => {
                this.getDimensions();
            });
    }
}