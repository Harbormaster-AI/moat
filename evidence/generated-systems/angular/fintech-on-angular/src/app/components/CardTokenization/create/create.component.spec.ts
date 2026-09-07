
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCardTokenizationComponent } from './create.component';
import { CardTokenizationService } from '../../../services/CardTokenization.service';
import { Router } from '@angular/router';

describe('CreateCardTokenizationComponent', () => {
  let component: CreateCardTokenizationComponent;
  let fixture: ComponentFixture<CreateCardTokenizationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCardTokenizationComponent
      ],
      providers: [
        CardTokenizationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCardTokenizationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});