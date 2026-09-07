
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PositionService } from '../../../services/Position.service';
import { Position } from '../../../models/Position';

@Component({
    selector: 'app-index-position',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPositionComponent implements OnInit {

    positions: Position[] = [];

    constructor(
        private router: Router,
        private service: PositionService
) {}

    ngOnInit(): void {
        this.getPositions();
}

    getPositions(): void {
        this.service.getPositions().subscribe((res) => {
        this.positions = res;
    });
}

    deletePosition(id: any): void {
        this.service.deletePosition(id)
            .subscribe(() => {
                this.getPositions();
            });
    }
}