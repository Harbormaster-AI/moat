
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateMatterComponent } from './create.component';
import { MatterService } from '../../../services/Matter.service';
import { Router } from '@angular/router';

describe('CreateMatterComponent', () => {
  let component: CreateMatterComponent;
  let fixture: ComponentFixture<CreateMatterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateMatterComponent
      ],
      providers: [
        MatterService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateMatterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});