
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { EncounterService } from '../../../services/Encounter.service';
import { Encounter } from '../../../models/Encounter';

@Component({
    selector: 'app-index-encounter',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexEncounterComponent implements OnInit {

    encounters: Encounter[] = [];

    constructor(
        private router: Router,
        private service: EncounterService
) {}

    ngOnInit(): void {
        this.getEncounters();
}

    getEncounters(): void {
        this.service.getEncounters().subscribe((res) => {
        this.encounters = res;
    });
}

    deleteEncounter(id: any): void {
        this.service.deleteEncounter(id)
            .subscribe(() => {
                this.getEncounters();
            });
    }
}