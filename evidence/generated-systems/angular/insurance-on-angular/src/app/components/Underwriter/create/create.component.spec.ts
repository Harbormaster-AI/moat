
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateUnderwriterComponent } from './create.component';
import { UnderwriterService } from '../../../services/Underwriter.service';
import { Router } from '@angular/router';

describe('CreateUnderwriterComponent', () => {
  let component: CreateUnderwriterComponent;
  let fixture: ComponentFixture<CreateUnderwriterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateUnderwriterComponent
      ],
      providers: [
        UnderwriterService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateUnderwriterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});