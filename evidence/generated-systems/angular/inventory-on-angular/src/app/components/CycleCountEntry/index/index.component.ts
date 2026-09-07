
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CycleCountEntryService } from '../../../services/CycleCountEntry.service';
import { CycleCountEntry } from '../../../models/CycleCountEntry';

@Component({
    selector: 'app-index-cycleCountEntry',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCycleCountEntryComponent implements OnInit {

    cycleCountEntrys: CycleCountEntry[] = [];

    constructor(
        private router: Router,
        private service: CycleCountEntryService
) {}

    ngOnInit(): void {
        this.getCycleCountEntrys();
}

    getCycleCountEntrys(): void {
        this.service.getCycleCountEntrys().subscribe((res) => {
        this.cycleCountEntrys = res;
    });
}

    deleteCycleCountEntry(id: any): void {
        this.service.deleteCycleCountEntry(id)
            .subscribe(() => {
                this.getCycleCountEntrys();
            });
    }
}