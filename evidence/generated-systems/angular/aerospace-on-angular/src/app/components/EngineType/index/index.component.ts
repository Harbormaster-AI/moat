
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { EngineTypeService } from '../../../services/EngineType.service';
import { EngineType } from '../../../models/EngineType';

@Component({
    selector: 'app-index-engineType',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexEngineTypeComponent implements OnInit {

    engineTypes: EngineType[] = [];

    constructor(
        private router: Router,
        private service: EngineTypeService
) {}

    ngOnInit(): void {
        this.getEngineTypes();
}

    getEngineTypes(): void {
        this.service.getEngineTypes().subscribe((res) => {
        this.engineTypes = res;
    });
}

    deleteEngineType(id: any): void {
        this.service.deleteEngineType(id)
            .subscribe(() => {
                this.getEngineTypes();
            });
    }
}