
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAdjusterComponent } from './create.component';
import { AdjusterService } from '../../../services/Adjuster.service';
import { Router } from '@angular/router';

describe('CreateAdjusterComponent', () => {
  let component: CreateAdjusterComponent;
  let fixture: ComponentFixture<CreateAdjusterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAdjusterComponent
      ],
      providers: [
        AdjusterService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAdjusterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});