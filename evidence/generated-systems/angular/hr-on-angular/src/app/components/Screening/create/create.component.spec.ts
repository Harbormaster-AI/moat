
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateScreeningComponent } from './create.component';
import { ScreeningService } from '../../../services/Screening.service';
import { Router } from '@angular/router';

describe('CreateScreeningComponent', () => {
  let component: CreateScreeningComponent;
  let fixture: ComponentFixture<CreateScreeningComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateScreeningComponent
      ],
      providers: [
        ScreeningService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateScreeningComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});