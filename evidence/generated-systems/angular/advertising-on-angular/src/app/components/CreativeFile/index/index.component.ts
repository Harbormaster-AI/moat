
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CreativeFileService } from '../../../services/CreativeFile.service';
import { CreativeFile } from '../../../models/CreativeFile';

@Component({
    selector: 'app-index-creativeFile',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCreativeFileComponent implements OnInit {

    creativeFiles: CreativeFile[] = [];

    constructor(
        private router: Router,
        private service: CreativeFileService
) {}

    ngOnInit(): void {
        this.getCreativeFiles();
}

    getCreativeFiles(): void {
        this.service.getCreativeFiles().subscribe((res) => {
        this.creativeFiles = res;
    });
}

    deleteCreativeFile(id: any): void {
        this.service.deleteCreativeFile(id)
            .subscribe(() => {
                this.getCreativeFiles();
            });
    }
}