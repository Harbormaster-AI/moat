
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateInspectionResultComponent } from './create.component';
import { InspectionResultService } from '../../../services/InspectionResult.service';
import { Router } from '@angular/router';

describe('CreateInspectionResultComponent', () => {
  let component: CreateInspectionResultComponent;
  let fixture: ComponentFixture<CreateInspectionResultComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateInspectionResultComponent
      ],
      providers: [
        InspectionResultService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateInspectionResultComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});