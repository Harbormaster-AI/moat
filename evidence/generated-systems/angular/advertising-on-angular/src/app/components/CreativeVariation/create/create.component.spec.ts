
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCreativeVariationComponent } from './create.component';
import { CreativeVariationService } from '../../../services/CreativeVariation.service';
import { Router } from '@angular/router';

describe('CreateCreativeVariationComponent', () => {
  let component: CreateCreativeVariationComponent;
  let fixture: ComponentFixture<CreateCreativeVariationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCreativeVariationComponent
      ],
      providers: [
        CreativeVariationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCreativeVariationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});