
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditPolicyCoverageComponent } from './edit.component';
import { PolicyCoverageService } from '../../../services/PolicyCoverage.service';

describe('EditPolicyCoverageComponent', () => {
  let component: EditPolicyCoverageComponent;
  let fixture: ComponentFixture<EditPolicyCoverageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditPolicyCoverageComponent
      ],
      providers: [
        PolicyCoverageService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditPolicyCoverageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});