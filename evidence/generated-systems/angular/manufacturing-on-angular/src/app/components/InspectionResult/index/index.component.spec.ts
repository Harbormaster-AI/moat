
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInspectionResultComponent } from './index.component';
import { InspectionResultService } from '../../../services/InspectionResult.service';

describe('IndexInspectionResultComponent', () => {
  let component: IndexInspectionResultComponent;
  let fixture: ComponentFixture<IndexInspectionResultComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInspectionResultComponent
      ],
      providers: [
        InspectionResultService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInspectionResultComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});