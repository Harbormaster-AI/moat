
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePolicyCoverageComponent } from './create.component';
import { PolicyCoverageService } from '../../../services/PolicyCoverage.service';
import { Router } from '@angular/router';

describe('CreatePolicyCoverageComponent', () => {
  let component: CreatePolicyCoverageComponent;
  let fixture: ComponentFixture<CreatePolicyCoverageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePolicyCoverageComponent
      ],
      providers: [
        PolicyCoverageService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePolicyCoverageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});