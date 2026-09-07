
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCreativeVariationComponent } from './index.component';
import { CreativeVariationService } from '../../../services/CreativeVariation.service';

describe('IndexCreativeVariationComponent', () => {
  let component: IndexCreativeVariationComponent;
  let fixture: ComponentFixture<IndexCreativeVariationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCreativeVariationComponent
      ],
      providers: [
        CreativeVariationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCreativeVariationComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});