
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCompetencyComponent } from './create.component';
import { CompetencyService } from '../../../services/Competency.service';
import { Router } from '@angular/router';

describe('CreateCompetencyComponent', () => {
  let component: CreateCompetencyComponent;
  let fixture: ComponentFixture<CreateCompetencyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCompetencyComponent
      ],
      providers: [
        CompetencyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCompetencyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});