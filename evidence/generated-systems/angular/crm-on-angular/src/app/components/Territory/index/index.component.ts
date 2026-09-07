
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TerritoryService } from '../../../services/Territory.service';
import { Territory } from '../../../models/Territory';

@Component({
    selector: 'app-index-territory',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexTerritoryComponent implements OnInit {

    territorys: Territory[] = [];

    constructor(
        private router: Router,
        private service: TerritoryService
) {}

    ngOnInit(): void {
        this.getTerritorys();
}

    getTerritorys(): void {
        this.service.getTerritorys().subscribe((res) => {
        this.territorys = res;
    });
}

    deleteTerritory(id: any): void {
        this.service.deleteTerritory(id)
            .subscribe(() => {
                this.getTerritorys();
            });
    }
}