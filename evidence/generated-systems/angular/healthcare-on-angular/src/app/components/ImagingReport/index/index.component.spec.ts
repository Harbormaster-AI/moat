
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexImagingReportComponent } from './index.component';
import { ImagingReportService } from '../../../services/ImagingReport.service';

describe('IndexImagingReportComponent', () => {
  let component: IndexImagingReportComponent;
  let fixture: ComponentFixture<IndexImagingReportComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexImagingReportComponent
      ],
      providers: [
        ImagingReportService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexImagingReportComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});