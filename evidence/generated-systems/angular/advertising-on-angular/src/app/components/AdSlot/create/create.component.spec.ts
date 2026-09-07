
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAdSlotComponent } from './create.component';
import { AdSlotService } from '../../../services/AdSlot.service';
import { Router } from '@angular/router';

describe('CreateAdSlotComponent', () => {
  let component: CreateAdSlotComponent;
  let fixture: ComponentFixture<CreateAdSlotComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAdSlotComponent
      ],
      providers: [
        AdSlotService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAdSlotComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});