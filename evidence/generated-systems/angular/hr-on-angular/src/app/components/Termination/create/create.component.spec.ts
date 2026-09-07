
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateTerminationComponent } from './create.component';
import { TerminationService } from '../../../services/Termination.service';
import { Router } from '@angular/router';

describe('CreateTerminationComponent', () => {
  let component: CreateTerminationComponent;
  let fixture: ComponentFixture<CreateTerminationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateTerminationComponent
      ],
      providers: [
        TerminationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateTerminationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});